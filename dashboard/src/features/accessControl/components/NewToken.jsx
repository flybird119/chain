import React from 'react'
import { BaseNew, FormContainer, FormSection, TextField, SelectField } from 'features/shared/components'
import { policyOptions } from 'features/accessControl/constants'
import { reduxForm } from 'redux-form'

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
          <TextField title='Token Name' fieldProps={guard_data.id} />
          <SelectField options={policyOptions}
            title='Policy'
            hint='Available policies are:

* `client-readwrite`: full access to the Client API
* `client-readonly`: access to read-only Client endpoints
* `network`: access to the Network API
* `monitoring`: access to monitoring-specific endpoints
* `internal`: access to multi-process synchronization endpoints (Raft, etc.)'
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

export default BaseNew.connect(
  BaseNew.mapStateToProps('accessControl'),
  BaseNew.mapDispatchToProps('accessControl'),
  reduxForm({
    form: 'newAccessGrantForm',
    fields,
  })(NewToken)
)
