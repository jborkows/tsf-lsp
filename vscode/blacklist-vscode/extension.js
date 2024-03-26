const { LanguageClient } = require("vscode-languageclient");
const vscode = require("vscode");

module.exports = {
  activate(context) {
    const executable = {
      command: "/home/jborkows/projects/tsf-lsp/tmp/main",
    };

    const serverOptions = {
      run: executable,
      debug: executable,
    };

    const clientOptions = {
      documentSelector: [
        {
          scheme: "file",
          language: "plaintext",
          pattern: "**/*.tsf",
        },
      ],
    };

    const client = new LanguageClient(
      "blacklist-extension-id",
      "Blacklister",
      serverOptions,
      clientOptions,
    );

    let disposable = vscode.commands.registerCommand(
      "some_command",
      function (arg) {
        // Handle the command execution
        vscode.window.showInformationMessage(
          `Command executed with argument: ${arg} ${JSON.stringify(vscode.workspace.workspaceFolders)}`,
        );
        
        client.sendRequest("workspace/executeCommand", {
          command: "some_comand",
          arguments: `${vscode.workspace.workspaceFolders[0].uri.fsPath}`,
        });
      },
    );

    context.subscriptions.push(disposable);
    context.subscriptions.push(client.start());
  },
};
