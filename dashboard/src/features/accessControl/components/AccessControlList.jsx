import React from 'react'
import AccessGrantListItem from './AccessGrantListItem'
import { connect } from 'react-redux'
import { TableList, PageTitle, PageContent } from 'features/shared/components'
import { replace } from 'react-router-redux'

import { actions } from 'features/accessControl'
import styles from './AccessControlList.scss'

class AccessControlList extends React.Component {
  render() {
    const tableListProps = {
      titles: ['ID', 'Policy']
    }

    const tokenList = <TableList {...tableListProps}>
      {this.props.tokens.map(item => <AccessGrantListItem key={item.id} item={item} />)}
    </TableList>

    const certList = <TableList {...tableListProps}>
      {this.props.certs.map(item => <AccessGrantListItem key={item.id} item={item} />)}
    </TableList>

    return (<div>
      <PageTitle title='Access Control' />

      <PageContent>
        <div className={`btn-group ${styles.btnGroup}`} role='group'>
          <button
            className={`btn btn-default ${styles.btn} ${this.props.tokensButtonStyle}`}
            onClick={this.props.showTokens}>
              Tokens
          </button>

          <button
            className={`btn btn-default ${styles.btn} ${this.props.certificatesButtonStyle}`}
            onClick={this.props.showCertificates}>
              Certificates
          </button>
        </div>

        {this.props.tokensSelected && <div>
          <button
            key='showCreate'
            className={`btn btn-primary ${styles.newBtn}`}
            onClick={this.props.showTokenCreate}>
              + New access grant
          </button>

          {tokenList}
        </div>}

        {this.props.certificatesSelected && certList}
      </PageContent>
    </div>)
  }
}

const mapStateToProps = (state, ownProps) => {
  const items = Object.values(state.accessControl.items)
  const tokensSelected = ownProps.location.query.type == 'token'
  const certificatesSelected = ownProps.location.query.type != 'token'

  return {
    tokens: items.filter(item => item.type == 'access_token'),
    certs: items.filter(item => item.type == 'x509'),
    tokensSelected,
    certificatesSelected,
    tokensButtonStyle: tokensSelected && styles.active,
    certificatesButtonStyle: certificatesSelected && styles.active,
  }
}

const mapDispatchToProps = (dispatch) => ({
  showTokens: () => dispatch(replace('/access_control?type=token')),
  showCertificates: () => dispatch(replace('/access_control?type=certificate')),
  showTokenCreate: () => dispatch(actions.showTokenCreate),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AccessControlList)
