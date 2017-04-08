import { chainClient } from 'utility/environment'
import { baseCreateActions, baseUpdateActions, baseListActions } from 'features/shared/actions'

const type = 'account'

const list = baseListActions(type, { defaultKey: 'alias' })
const form = baseCreateActions(type, {
  jsonFields: ['tags'],
  intFields: ['quorum'],
  redirectToShow: true,
})
const update = baseUpdateActions(type, {
  jsonFields: ['tags'],
  redirectToShow: true,
})

let actions = {
  ...list,
  ...form,
  ...update,
  createReceiver: (data) => () => {
    return chainClient().accounts.createReceiver(data)
  }
}

export default actions
