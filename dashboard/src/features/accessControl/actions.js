import { chainClient } from 'utility/environment'
import { baseListActions } from 'features/shared/actions'

export default baseListActions('accessControl', {
  clientApi: () => chainClient().accessControl
})

// fetchItems,
// fetchPage,
// fetchAll,
// deleteItem,
// pushList,
// didLoadAutocomplete
