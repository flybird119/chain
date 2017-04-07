import { combineReducers } from 'redux'
import { reducers } from 'features/shared'

const type = 'accessControl'

export default combineReducers({
  items: reducers.itemsReducer(type),
  queries: reducers.queriesReducer(type),
})

// items
// queries
