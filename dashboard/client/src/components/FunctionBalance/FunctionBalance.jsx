import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { Progress } from 'reactstrap';
import { useMonetizationCounter } from 'react-web-monetization'

import { functionsApi } from '../../api/functionsApi';

const FunctionBalance = ({ fn }) => {
  const { receipt } = useMonetizationCounter();
  const [remainingInvocations, setRemainingInvocations] = useState(0);

  const getBalance = async () => {
    const res = await functionsApi.fetchFunctionBalance(fn.name);
    setRemainingInvocations(Math.min(100, res.remainingInvocations));
  }

  const submitReceipt = async (receipt) => {
    await fetch(window.RECEIPTS_URL, {
      method: 'POST',
      body: receipt
    });
    getBalance()
  }

  useEffect(() => {
    getBalance()
  }, [])

  useEffect(() => {
    if (receipt !== null) {
      submitReceipt(receipt)
    }
  }, [receipt])

  return (
    <div>
      <Progress multi={true} className="mt-3 d-flex justify-content-center">
        <Progress bar={true} color="success" value={remainingInvocations} />
        <Progress bar={true} color="danger" value={100 - remainingInvocations} />
      </Progress>
      <span className="font-weight-bold">{remainingInvocations >= 100 ? '100+' : remainingInvocations}</span> invocations remaining
    </div>
  );
};

FunctionBalance.propTypes = {
  fn: PropTypes.object.isRequired
};

export { FunctionBalance };
