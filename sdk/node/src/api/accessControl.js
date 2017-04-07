const shared = require('../shared')

const accessControl = (/* client */) => ({
  create: (/* params , cb */) =>
    Promise.resolve({message: 'ok'}),
    // shared.create(client, '/acl-grant', params, {skipArray: true}, cb),

  delete: (params, cb) => shared.tryCallback(
    () => Promise.resolve({message: 'ok'}),
    // client.request('/acl-revoke', params),
    cb
  ),

  query: (/* params, cb */) =>
    Promise.resolve({
      items: [
        {
          id: 'domtoken',
          type: 'access_token',
          policy: 'client-readwrite',
        },
        {
          id: 'SN: visa.com',
          type: 'x509',
          policy: 'network',
        },
      ],
      next: {},
      lastPage: true,
    })
    // shared.query(client, 'accessTokens', '/list-acl-grants', params, {cb}),
})

module.exports = accessControl
