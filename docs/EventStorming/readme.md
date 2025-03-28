# Event Storming

In this section we declare all possible event in application. Creating ubiquitous language from each domain.

## Ledger
Service ini bertugas untuk mencatat transaksi yang dilakukan oleh merchant.

### Events
- Payout Created 
  - This event when merchant maker create new Payouts. Merchant can create single or bulk Payout
- Payout Approved
	- This event when merchant approval approve the Payouts.
- Payout Rejected
	- This event when merchant approval reject the Payouts. Approval merchant need add reason when reject Payouts.
- Payout Started
	- This event when system starting the Payout.
- Payout Success
	- This event when service got callback Payout already succeed.
- Payout Pending
	- This event when service got callback Payout in pending state.
- Payout Failed
	- This event when service got callback orchestrator failed to process Payout.
- Payin Created
	- This event when merchant create Payin method.
- Payin Cancelled
	- This event when merchant cancell the Payin method, like when they change the payment method.
- Account Inquiried
	- This event when merchant checking beneficiary account.

## Orchestrator
Servis ini bertanggung jawab untuk melakukan trigger transfer (payout) dan menerima payment (payin). Dalam hal Payout, service ini juga bertugas melakukan routing antar processor (Banks, other payment gateway).

### Events
- Transfer Started
  - This event when transfer routing started.
- Transfer Rerouted
  - This event when transfer got response failed from processor, and need to reroute.
- Transfer Finished
  - This event when transfer got final state, or pending state. Service will send callback to Ledger to update payout status.
- Payment Success
	- This event when orchestrator got payment callback with status success/paid.
- Payment Expired
	- This event when Payment not paid until expired date pass.
- Payment Status Checked
  - This event when service check status of payment periodically.

## Snap Processor
Service ini bertugas untuk mentranslate payload request dari/ke [Standard Snap](https://apidevportal.aspi-indonesia.or.id/).

### Events
- Transfer Triggered
  - This event when beneficiary account inquiried to third party.
- Account Inquiried
  - This event when transfer executed to third party.
- Transfer Status Checked
  - This event when transfer status checked to third party.
- Virtual Account Created
  - This event when va is created.
- Virtual Account Inquiried
  - This event when va is inquiried from third party.
- Virtual Account Paid
  - This event when service got va payment callback from third party.
- Virtual Account Status Checked
  - This event when service check status payment to third party.
- QRIS Created
  - This event when qris is created.
- QRIS Paid
  - This event when service got qris payment callback from third party.
- QRIS Status Checked
  - This event when service check status payment to third party.

## Notification
### Events
- Callback Registered
  - This event when merchant register their callback endpoint.
- Callback Sended
  - This event when service success send callback.
- Callback Failed
  - This event when service failed to send callback to merchant service.
- Notification Sended
  - This event when service success send notification.