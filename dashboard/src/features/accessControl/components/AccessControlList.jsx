import React from 'react'
import AccessGrantListItem from './AccessGrantListItem'
import { connect } from 'react-redux'
import { TableList, PageTitle, PageContent } from 'features/shared/components'

// export default BaseList.connect(
//   BaseList.mapStateToProps(type, AccessGrantListItem, {
//     skipQuery: true,
//     label: 'Access control',
//     wrapperComponent: TableList,
//     wrapperProps: {
//       titles: ['ID', 'Grant']
//     }
//   }),
//   BaseList.mapDispatchToProps(type)
// )

class AccessControlList extends React.Component {
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
        {tokenList}
        {certList}
      </PageContent>
    </div>)
  }
}

const mapStateToProps = (state) => {
  const items = Object.values(state.accessControl.items)
  return {
    tokens: items.filter(item => item.type == 'access_token'),
    certs: items.filter(item => item.type == 'x509'),
  }
}

const mapDispatchToProps = ( ) => ({})


export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AccessControlList)
