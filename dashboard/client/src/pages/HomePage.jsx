import React, { Component } from 'react';
import { FunctionEmptyState } from "../components/FunctionEmptyState";
import {
  Card,
  CardHeader,
} from 'reactstrap';

export class HomePage extends Component {
  render() {
    return (
      <Card outline color="success">
        <CardHeader className="bg-success color-success">
          GitHub App: <a
            href={window.GITHUB_APP_URL}
            target="_blank"
          >
            {window.GITHUB_APP_URL}
          </a>
        </CardHeader>
        <FunctionEmptyState />
      </Card>
    );
  }
}
