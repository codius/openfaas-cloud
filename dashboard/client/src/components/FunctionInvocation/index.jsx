import React from 'react';
import {
  Badge,
  ListGroup,
  ListGroupItem,
  Nav,
  NavLink
} from 'reactstrap';

import './FunctionInvocation.css';
import { FunctionBalance } from '../FunctionBalance';

const OPTIONS = {
  '1hr': '60m',
  '24hr': '1440m'
};

export class FunctionInvocation extends React.Component {
  state = {
    selected: '1hr'
  };

  render() {
    const { functionInvocationData } = this.props;
    let { success, failure } = functionInvocationData;
    const navLinks = Object.keys(OPTIONS).map(option => {
      return (
        <NavLink
          key={option}
          href="#"
          active={option === this.state.selected}
          onClick={() => this.navLinkClickHandle(option)}
        >
          {option}
        </NavLink>
      );
    });

    return (
      <div className="">
        <FunctionBalance fn={this.props.fn} />
        <Nav className="d-flex justify-content-center">
          <span className="d-flex align-items-center mr-4 font-weight-bold">
            Period:
          </span>
          {navLinks}
        </Nav>
        <div className="mt-3 mx-1 row flex-row border">
          <div className="d-flex col-6 flex-column align-items-center border-right p-2">
            <h5 className="mt-1">
              <Badge color="success">{success}</Badge>
            </h5>
            <span>Success</span>
          </div>
          <div className="d-flex col-6 flex-column align-items-center p-2">
            <h5 className="mt-1">
              <Badge color="danger">{failure}</Badge>
            </h5>
            <span>Error</span>
          </div>
        </div>
      </div>
    );
  }

  navLinkClickHandle = option => {
    this.setState({
      selected: option
    });
    this.props.changeFunctionInvocationTimePeriod(OPTIONS[option]);
  };
}
