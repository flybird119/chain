const shared = require('../shared')

const accessControl = (client) => ({
  create: (params, cb) =>
    shared.create(client, '/acl-grant', params, {skipArray: true}, cb),

  delete: (params, cb) => shared.tryCallback(
    client.request('/acl-revoke', params),
    cb
  ),

  query: (params, cb) =>
    shared.query(client, 'accessTokens', '/list-acl-grants', params, {cb}),
})

module.exports = accessControl
