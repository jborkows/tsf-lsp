
  const { LanguageClient } = require('vscode-languageclient')

module.exports = {
  activate(context) {
    const executable = {
      command: '/home/jborkows/projects/tsf-lsp/tmp/main',
    }

    const serverOptions = {
      run: executable,
      debug: executable,
    }

    const clientOptions = {
      documentSelector: [{
        scheme: 'file',
        language: 'plaintext',
        pattern: "**/*.tsf"
      }],
    }

    const client = new LanguageClient(
      'blacklist-extension-id',
      'Blacklister',
      serverOptions,
      clientOptions
    )

    context.subscriptions.push(client.start())
  },

}
