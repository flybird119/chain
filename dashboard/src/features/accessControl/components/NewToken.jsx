import React from 'react'
import { BaseNew, FormContainer, FormSection, TextField, SelectField } from 'features/shared/components'
import { policyOptions } from 'features/accessControl/constants'
import { reduxForm } from 'redux-form'
import { actions } from 'features/accessControl'

class NewToken extends React.Component {
  render() {
    const {
      fields: { guard_data, policy },
      error,
      handleSubmit,
      submitting
    } = this.props

    return(
      <FormContainer
        error={error}
        label='New access token'
        onSubmit={handleSubmit(this.props.submitForm)}
        submitting={submitting} >

        <FormSection title='Token information'>
          <TextField title='Token Name' fieldProps={guard_data.id} autoFocus={true} />
          <SelectField options={policyOptions}
            title='Policy'
            hint='Available policies are:
* `client-readwrite`: full access to the Client API
* `client-readonly`: access to read-only Client endpoints
* `network`: access to the Network API
* `monitoring`: access to monitoring-specific endpoints'
            fieldProps={policy} />
        </FormSection>

      </FormContainer>
    )
  }
}

const fields = [
  'guard_type',
  'guard_data.id',
  'policy',
]

const validate = values => {
  const errors = {}

  if (!values.policy) {
    errors.policy = 'Policy is required'
  }
  if (!values.guard_data.id) {
    errors.guard_data = {id: 'Token name is required'}
  }

  return errors
}

const mapDispatchToProps = (dispatch) => ({
  submitForm: (data) => dispatch(actions.submitTokenForm(data))
})

export default BaseNew.connect(
  BaseNew.mapStateToProps('accessControl'),
  mapDispatchToProps,
  reduxForm({
    form: 'newAccessGrantForm',
    fields,
    validate
  })(NewToken)
)
