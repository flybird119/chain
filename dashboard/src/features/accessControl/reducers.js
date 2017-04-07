import { combineReducers } from 'redux'
import { reducers } from 'features/shared'

const type = 'accessControl'
const idFunc = (item, index) => index

const itemsReducer = (state = {}, action) => {
  if (action.type == 'APPEND_ACCESSCONTROL_PAGE') {
    const newState = {}
    action.param.items.forEach((item, index) => {
      item.id = `acl-${index}`
      newState[index] = item
    })
    return newState
  }
  return state
}

const listViewReducer = combineReducers({
  itemIds: reducers.queryItemsReducer(type, idFunc),
  cursor: reducers.queryCursorReducer(type),
  queryTime: reducers.queryTimeReducer(type),
})

const queriesReducer = (state = {}, action) => {
  if (action.type == 'APPEND_ACCESSCONTROL_PAGE') {
    const query = action.param.next.filter || ''
    const list = state[query] || {}

    return {
      [`${query}`]: listViewReducer(list, action)
    }
  }

  return state
}


export default combineReducers({
  items: itemsReducer,
  queries: queriesReducer
})
