import psutil
import subprocess
import win32gui, win32process, win32con

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

    close_windows(process)
    try:
        process.wait(timeout=10)
    except psutil.TimeoutExpired:
        stop_process(process)

    return subprocess.Popen(cmdline, cwd=cwd, env=env)

def set_username(cmdline, username):
    return cmdline

def main():
    process = find_minecraft_process()
    if process is None:
        print('Could not find Minecraft. Is it running?')
        return
    
    username = input('Enter your fake username: ')
    cmdline = set_username(process.cmdline(), username)

    print('Restarting Minecraft...')
    new_process = restart_process(process, cmdline)
    pass

if __name__ == '__main__':
    main()
