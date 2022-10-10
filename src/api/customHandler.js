const { waitAndPrint } = require('../helper/delayed');
const { sendArgs } = require('/opt/proxy');

const handler = async (event, context) => {
  await sendArgs(event, context);
  await waitAndPrint({
    prefix: 'customHandler',
    timeout: 500,
  });
  return {
    message: 'Done doing something',
  };
};

module.exports = { handler };
