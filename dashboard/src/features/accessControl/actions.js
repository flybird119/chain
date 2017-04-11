import { chainClient } from 'utility/environment'
import { baseListActions } from 'features/shared/actions'
import { push } from 'react-router-redux'

export default {
  ...baseListActions('accessControl', {
    clientApi: () => chainClient().accessControl
  }),
  showTokenCreate: push('access_control/create-token'),
  submitTokenForm: data => {
    const body = {...data}

    body.guard_type = 'access_token'

    return function(dispatch) {
      return chainClient().accessControl.create(body).then(resp => {
        dispatch({ type: 'CREATED_ACCESSTOKEN', resp })
        dispatch(push({pathname: 'access_control', state: {preserveFlash: true}}))
      }, err => Promise.reject(err))
    }
  },
  showAddCertificate: push('access_control/add-certificate'),
  submitCertificateForm: data => {
    const body = {...data}

    body.guard_type = 'x509'

    return function(dispatch) {
      return chainClient().accessControl.create(body).then(resp => {
        dispatch({ type: 'CREATED_ACCESSX509', resp })
        dispatch(push({pathname: 'access_control?type=certificate', state: {preserveFlash: true}}))
      }, err => Promise.reject(err))
    }

  },
}

// fetchItems,
// fetchPage,
// fetchAll,
// deleteItem,
// pushList,
// x didLoadAutocomplete,
// √ showcreate,
// created,
// √ submitForm
