import psutil
import os, subprocess
import win32gui, win32process, win32con

CLOSE_EXISTING_INSTANCE = True

def find_minecraft_process():
    for process in psutil.process_iter(attrs=['pid', 'name']):
        try:
            if process.info['name'].lower() == 'javaw.exe' and 'net.minecraft.client.main.Main' in process.cmdline():
                return process
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass
    
    return None

def close_windows(process):
    def callback(hwnd, _):
        _, window_pid = win32process.GetWindowThreadProcessId(hwnd)
        if window_pid == process.pid:
            win32gui.PostMessage(hwnd, win32con.WM_CLOSE, 0, 0)
    win32gui.EnumWindows(callback, None)

def stop_process(process):
    try:
        process.terminate()
        process.wait(timeout=10)
    except psutil.TimeoutExpired:
        process.kill()

def restart_process(process: psutil.Process, cmdline: list[str]) -> psutil.Process:
    cwd = process.cwd()
    env = process.environ()

    if CLOSE_EXISTING_INSTANCE:
        print('Closing Minecraft...')
        close_windows(process)
        try:
            process.wait(timeout=30)
        except psutil.TimeoutExpired:
            stop_process(process)

    print('Restarting Minecraft...' if CLOSE_EXISTING_INSTANCE else 'Opening Minecraft...')
    subprocess.Popen(cmdline, cwd=cwd, env=env, creationflags=subprocess.CREATE_NEW_CONSOLE)

def set_username(cmdline, username):
    try:
        i = cmdline.index('--username')
        cmdline[i+1] = username
        return cmdline
    except ValueError:
        return None

def main():
    process = find_minecraft_process()
    if process is None:
        print('Could not find Minecraft: Java Edition. Is it running?')
        return
    
    username = input('Enter your fake username: ')
    cmdline = process.cmdline()
    modified_cmdline = set_username(process.cmdline(), username)
    if modified_cmdline is None:
        with open('report.txt', 'a') as report:
            report.writelines(arg + '\n' for arg in cmdline)
        report_path = os.path.abspath('report.txt')
        print('Could not figure out how to change the username. This script likely needs to be updated. Please send the developer a copy of the report located here: ' + report_path)
        return

    restart_process(process, modified_cmdline)
    print('Done! Minecraft may take a moment to finish opening.')

if __name__ == '__main__':
    main()
    os.system('pause')
