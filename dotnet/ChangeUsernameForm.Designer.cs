namespace dotnet;

partial class ChangeUsernameForm
{
    /// <summary>
    ///  Required designer variable.
    /// </summary>
    private System.ComponentModel.IContainer components = null;

    /// <summary>
    ///  Clean up any resources being used.
    /// </summary>
    /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
    protected override void Dispose(bool disposing)
    {
        if (disposing && (components != null))
        {
            components.Dispose();
        }
        base.Dispose(disposing);
    }

    #region Windows Form Designer generated code

    /// <summary>
    ///  Required method for Designer support - do not modify
    ///  the contents of this method with the code editor.
    /// </summary>
    private void InitializeComponent()
    {
        mainTableLayoutPanel = new TableLayoutPanel();
        instanceListBox = new ListBox();
        usernameLabel = new Label();
        usernameTextBox = new TextBox();
        applyButton = new Button();
        mainTableLayoutPanel.SuspendLayout();
        SuspendLayout();
        // 
        // mainTableLayoutPanel
        // 
        mainTableLayoutPanel.Anchor = AnchorStyles.Top | AnchorStyles.Bottom | AnchorStyles.Left | AnchorStyles.Right;
        mainTableLayoutPanel.ColumnCount = 2;
        mainTableLayoutPanel.ColumnStyles.Add(new ColumnStyle());
        mainTableLayoutPanel.ColumnStyles.Add(new ColumnStyle(SizeType.Percent, 100F));
        mainTableLayoutPanel.ColumnStyles.Add(new ColumnStyle(SizeType.Absolute, 20F));
        mainTableLayoutPanel.Controls.Add(instanceListBox, 0, 0);
        mainTableLayoutPanel.Controls.Add(usernameLabel, 0, 1);
        mainTableLayoutPanel.Controls.Add(usernameTextBox, 1, 1);
        mainTableLayoutPanel.Controls.Add(applyButton, 0, 2);
        mainTableLayoutPanel.Location = new Point(6, 6);
        mainTableLayoutPanel.Name = "mainTableLayoutPanel";
        mainTableLayoutPanel.RowCount = 3;
        mainTableLayoutPanel.RowStyles.Add(new RowStyle(SizeType.Percent, 100F));
        mainTableLayoutPanel.RowStyles.Add(new RowStyle());
        mainTableLayoutPanel.RowStyles.Add(new RowStyle());
        mainTableLayoutPanel.Size = new Size(446, 260);
        mainTableLayoutPanel.TabIndex = 0;
        // 
        // instanceListBox
        // 
        mainTableLayoutPanel.SetColumnSpan(instanceListBox, 2);
        instanceListBox.Dock = DockStyle.Fill;
        instanceListBox.FormattingEnabled = true;
        instanceListBox.ItemHeight = 20;
        instanceListBox.Location = new Point(3, 3);
        instanceListBox.Name = "instanceListBox";
        instanceListBox.Size = new Size(440, 186);
        instanceListBox.TabIndex = 0;
        // 
        // usernameLabel
        // 
        usernameLabel.AutoSize = true;
        usernameLabel.Dock = DockStyle.Left;
        usernameLabel.Location = new Point(3, 192);
        usernameLabel.Name = "usernameLabel";
        usernameLabel.Size = new Size(75, 33);
        usernameLabel.TabIndex = 1;
        usernameLabel.Text = "Username";
        usernameLabel.TextAlign = ContentAlignment.MiddleLeft;
        // 
        // usernameTextBox
        // 
        usernameTextBox.Dock = DockStyle.Top;
        usernameTextBox.Location = new Point(84, 195);
        usernameTextBox.Name = "usernameTextBox";
        usernameTextBox.Size = new Size(359, 27);
        usernameTextBox.TabIndex = 2;
        // 
        // applyButton
        // 
        mainTableLayoutPanel.SetColumnSpan(applyButton, 2);
        applyButton.Dock = DockStyle.Bottom;
        applyButton.Location = new Point(3, 228);
        applyButton.Name = "applyButton";
        applyButton.Size = new Size(440, 29);
        applyButton.TabIndex = 3;
        applyButton.Text = "Apply";
        applyButton.UseVisualStyleBackColor = true;
        // 
        // ChangeUsernameForm
        // 
        AutoScaleDimensions = new SizeF(8F, 20F);
        AutoScaleMode = AutoScaleMode.Font;
        ClientSize = new Size(458, 272);
        Controls.Add(mainTableLayoutPanel);
        MinimumSize = new Size(257, 174);
        Name = "ChangeUsernameForm";
        ShowIcon = false;
        Text = "Change Username";
        mainTableLayoutPanel.ResumeLayout(false);
        mainTableLayoutPanel.PerformLayout();
        ResumeLayout(false);
    }

    #endregion

    private TableLayoutPanel mainTableLayoutPanel;
    private ListBox instanceListBox;
    private Label usernameLabel;
    private TextBox usernameTextBox;
    private Button applyButton;
}
