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
          `Command executed with argument: ${arg}`,
        );
        console.log("Command 'some_command' executed with arguments:", arg);
        client.sendRequest("workspace/executeCommand", {
          command: "some_comand",
          arguments: arg,
        });
      },
    );

    context.subscriptions.push(disposable);
    context.subscriptions.push(client.start());
  },
};
