import { baseCreateActions, baseUpdateActions, baseListActions } from 'features/shared/actions'

const type = 'asset'

const list = baseListActions(type, { defaultKey: 'alias' })
const form = baseCreateActions(type, {
  jsonFields: ['tags', 'definition'],
  intFields: ['quorum'],
  redirectToShow: true,
})
const update = baseUpdateActions(type, {
  jsonFields: ['tags'],
  redirectToShow: true,
})

const actions = {
  ...list,
  ...form,
  ...update,
}
export default actions
