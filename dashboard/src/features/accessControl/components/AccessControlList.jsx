import React from 'react'
import { PageTitle, PageContent } from 'features/shared/components'

class AccessControlList extends React.Component {
  render() {
    return(
      <div>
        <PageTitle title='Access Control' />

        <PageContent>
          ACL goes here
        </PageContent>
      </div>
    )
  }
}

export default AccessControlList
