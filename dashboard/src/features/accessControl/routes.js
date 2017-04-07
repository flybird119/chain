import AccessControlList from './components/AccessControlList'
import { makeRoutes } from 'features/shared'

export default (store) => ({
  path: 'access_control',
  indexRoute: {
    component: AccessControlList
  }
})
// makeRoutes(store, 'access_control', AccessControlList, null, null)
