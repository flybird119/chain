import AccessGrantListItem from './AccessGrantListItem'
import { BaseList, TableList } from 'features/shared/components'

const type = 'accessControl'

export default BaseList.connect(
  BaseList.mapStateToProps(type, AccessGrantListItem, {
    wrapperComponent: TableList,
    wrapperProps: {
      titles: ['ID', 'Grant']
    }
  }),
  BaseList.mapDispatchToProps(type)
)
