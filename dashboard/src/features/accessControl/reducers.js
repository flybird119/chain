import { combineReducers } from 'redux'
import { reducers } from 'features/shared'

const type = 'accessControl'

const idFunc = (item) => JSON.stringify(item.guard_data)

export default combineReducers({
  items: reducers.itemsReducer(type, idFunc),
  queries: reducers.queriesReducer(type, idFunc),
})

// items
// queries
