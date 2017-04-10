import React from 'react'
import AccessGrantListItem from './AccessGrantListItem'
import { connect } from 'react-redux'
import { TableList, PageTitle, PageContent } from 'features/shared/components'
import styles from './AccessControlList.scss'

class AccessControlList extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      selected: 'tokens'
    }
  }

  render() {
    const tableListProps = {
      titles: ['ID', 'Policy']
    }

    const tokenList = <TableList {...tableListProps}>
      {this.props.tokens.map(item => <AccessGrantListItem item={item} />)}
    </TableList>

    const certList = <TableList {...tableListProps}>
      {this.props.certs.map(item => <AccessGrantListItem item={item} />)}
    </TableList>

    return (<div>
      <PageTitle title='Access Control' />

      <PageContent>
        <div className={`btn-group ${styles.btnGroup}`} role='group'>
          <button className={`btn btn-default ${'x'}`}>Tokens</button>
          <button className={`btn btn-default ${'x'}`}>Certificates</button>
        </div>

        {tokenList}
        {certList}
      </PageContent>
    </div>)
  }
}

const mapStateToProps = (state, ownProps) => {
  const items = Object.values(state.accessControl.items)
  return {
    tokens: items.filter(item => item.type == 'access_token'),
    certs: items.filter(item => item.type == 'x509'),
    selectedTab: ownProps.params.type,
  }
}

const mapDispatchToProps = ( ) => ({})


export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AccessControlList)
