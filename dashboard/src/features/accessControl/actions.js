import { chainClient } from 'utility/environment'
import { baseListActions } from 'features/shared/actions'
import { push } from 'react-router-redux'

export default {
  ...baseListActions('accessControl', {
    clientApi: () => chainClient().accessControl
  }),
  showTokenCreate: push('access_control/create/token')
}

// fetchItems,
// fetchPage,
// fetchAll,
// deleteItem,
// pushList,
// didLoadAutocomplete,
// showcreate,
// created,
// submitForm
