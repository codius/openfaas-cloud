import React, { useEffect } from 'react';
import { useMonetizationCounter } from 'react-web-monetization'

const ReceiptSubmitter = () => {
  const { receipt } = useMonetizationCounter();

  useEffect(() => {
    if (receipt !== null) {
      fetch(window.RECEIPTS_URL, {
        method: 'POST',
        body: receipt
      });
    }
  }, [receipt])

  return null;
};

export { ReceiptSubmitter };
