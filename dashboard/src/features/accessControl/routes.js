import AccessControlList from './components/AccessControlList'
import { makeRoutes } from 'features/shared'

export default (store) => makeRoutes(store, 'accessControl', AccessControlList, null, null, {
  path: 'access_control'
})
