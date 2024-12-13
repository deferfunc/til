const AWS = require('aws-sdk');
const axios = require('axios'); // HTTPリクエストを送信するためにaxiosを使用
const logger = console;
AWS.config.update({ region: process.env.AWS_DEFAULT_REGION });

// メールの件名をデコード
function decodeSubject(encodedSubject) {
  try {
    return decodeURIComponent(encodedSubject);
  } catch (e) {
    logger.info(`メールの件名のデコードに失敗しました: ${e}`);
    return encodedSubject;
  }
}

// SSM ParameterStore から指定されたパラメーター名で値を取得
async function getParameterStoreValue(parameterName) {
  const ssm = new AWS.SSM();
  try {
    const data = await ssm.getParameter({ Name: parameterName, WithDecryption: false }).promise();
    return data.Parameter.Value;
  } catch (err) {
    logger.error(`パラメーターの取得に失敗しました: ${err}`);
    throw err;
  }
}

// Bounce 処理
exports.handler = async (event, context) => {
  const messageJson = JSON.parse(event.Records[0].Sns.Message);
  const destinationEmail = messageJson.mail.destination[0];
  const slackWebhookUrl = await getParameterStoreValue(process.env.SLACK_WEBHOOK_URL);

  const subject = decodeSubject(messageJson.mail.commonHeaders.subject);
  const slackMessage = {
    text: (
      `以下のメールを送信しましたが、受信者に届きませんでした。\n\n` +
      `• メールアドレス: ${destinationEmail}\n` +
      `• メールの件名: ${subject}\n\n` +
      `本通知を受けて、お客様へのご確認をお願いします。\n\n`
    )
  };

  try {
    const response = await axios.post(slackWebhookUrl, slackMessage);
    logger.info(`Slackに通知が送信されました: ${response.status}`);
  } catch (err) {
    logger.error(`Slackへの通知に失敗しました: ${err}`);
  }
};
