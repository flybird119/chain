import React from 'react'
import { BaseUpdate, FormContainer, FormSection, JsonField, TextField } from 'features/shared/components'
import { reduxForm } from 'redux-form'

class Form extends React.Component {
  constructor(props) {
    super(props)

    this.submitWithErrors = this.submitWithErrors.bind(this)

    this.state = {}
  }

  submitWithErrors(data) {
    return new Promise((resolve, reject) => {
      this.props.submitForm(data, this.props.item.id)
        .catch((err) => reject({_error: err}))
    })
  }

  componentDidMount() {
    this.props.fetchItem(this.props.params.id).then(resp => {
      if (resp.items.length == 0) {
        this.setState({notFound: true})
      }
    })
  }

  render() {
    const {
      fields: { alias, tags },
      error,
      handleSubmit,
      submitting
    } = this.props

    return(
      <FormContainer
        error={error}
        label='Edit account'
        onSubmit={handleSubmit(this.submitWithErrors)}
        submitting={submitting} >

        <FormSection title='Account Information'>
          <TextField title='Alias' placeholder='Alias' fieldProps={alias} autoFocus={true} />
          <JsonField title='Tags' fieldProps={tags} />
        </FormSection>
      </FormContainer>
    )
  }
}

const validate = values => {
  const errors = {}

  const jsonFields = ['tags']
  jsonFields.forEach(key => {
    const fieldError = JsonField.validator(values[key])
    if (fieldError) { errors[key] = fieldError }
  })

  return errors
}

const fields = [
  'alias',
  'tags'
]

const mapStateToProps = (state, ownProps) => ({
  item: state.account.items[ownProps.params.id]
})

const initialValues = (state, ownProps) => {
  const item = state.account.items[ownProps.params.id]
  if (item) {
    const tags = Object.keys(item.tags).length === 0 ? '{\n\t\n}' : JSON.stringify(item.tags, null, 1)
    return {
      initialValues: {
        alias: item.alias,
        tags: tags
      }
    }
  }
  return {}
}

export default BaseUpdate.connect(
  mapStateToProps,
  BaseUpdate.mapDispatchToProps('account'),
  reduxForm({
    form: 'updateAccountForm',
    fields,
    validate
  }, initialValues)(Form)
)
