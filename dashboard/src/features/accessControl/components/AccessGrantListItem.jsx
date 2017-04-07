import React from 'react'

class ListItem extends React.Component {
  render() {
    const item = this.props.item
    return(
      <tr>
        <td>{item.id}</td>
        <td>{item.policy}</td>
        <td>Revoke</td>
      </tr>
    )
  }
}

export default ListItem
