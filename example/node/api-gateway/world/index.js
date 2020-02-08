module.exports.endpoint = (event, context, callback) => {
    const response = {
        statusCode: 200,
        body: JSON.stringify({
            message: `Simple example`,
        }),
    };

    callback(null, response);
};