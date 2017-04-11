import AccessControlList from './components/AccessControlList'
import NewToken from './components/NewToken'
import NewCertificate from './components/NewCertificate'
import { makeRoutes } from 'features/shared'

const checkParams = (nextState, replace) => {
  if (!['token', 'certificate'].includes(nextState.location.query.type)) {
    replace('/access_control?type=token')
  }
}

export default (store) => {
  const routes = makeRoutes(store, 'accessControl', AccessControlList, null, null, {
    path: 'access_control',
    name: 'Access control'
  })

  const existingOnEnter = routes.indexRoute.onEnter
  const existingOnChange = routes.indexRoute.onChange

  routes.indexRoute.onEnter = (nextState, replace) => {
    checkParams(nextState, replace)
    existingOnEnter(nextState, replace)
  }

  routes.indexRoute.onChange = (_, nextState, replace) => {
    checkParams(nextState, replace)
    existingOnChange(_, nextState, replace)
  }

  routes.childRoutes.push({
    path: 'create-token',
    component: NewToken
  })

  routes.childRoutes.push({
    path: 'add-certificate',
    component: NewCertificate
  })

  return routes
}
