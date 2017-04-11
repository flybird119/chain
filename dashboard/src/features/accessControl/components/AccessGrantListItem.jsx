import React from 'react'

class AccessGrantListItem extends React.Component {
  render() {
    const item = this.props.item
    let desc
    if (item.guard_type == 'access_token') {
      desc = item.guard_data.id
    } else {
      desc = <div>
        {Object.keys(item.guard_data).map(field => <p>
          {field}:
          <ul>
            {Object.keys(item.guard_data[field]).map(key => <li>
              {key}: {item.guard_data[field][key]}
            </li>)}
          </ul>
        </p>)}
      </div>
    }
    return(
      <tr>
        <td>{desc}</td>
        <td>{item.policy}</td>
        <td>
          <button className='btn btn-danger btn-xs' disabled>Revoke</button>
        </td>
      </tr>
    )
  }
}

export default AccessGrantListItem
