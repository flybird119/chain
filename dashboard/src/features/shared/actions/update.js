import { chainClient } from 'utility/environment'
import { push } from 'react-router-redux'

export default function(type, options = {}) {
  const listPath = options.listPath || `/${type}s`
  const updated = (param) => ({ type: `UPDATED_${type.toUpperCase()}`, param })

  return {
    updated,
    submitUpdateForm: (data, id) => {
      const clientApi = options.clientApi ? options.clientApi() : chainClient()[`${type}s`]
      let promise = Promise.resolve()

      if (typeof data.alias == 'string')  data.alias = data.alias.trim()

      return function(dispatch) {
        return promise.then(() => clientApi.update({
          id: id,
          newAlias: data.alias,
          tags: JSON.parse(data.tags),
        }).then((resp) => {
          dispatch(updated(resp))

          let postCreatePath = listPath
          if (options.redirectToShow) {
            postCreatePath = `${postCreatePath}/${resp.id}`
          }

          dispatch(push({
            pathname: postCreatePath,
            state: {
              preserveFlash: true
            }
          }))
        }))
      }
    }
  }
}
