import { Client } from "@heroiclabs/nakama-js";

(async () => {
    try {
        const client = new Client("defaultkey", "localhost", 7350);

        const deviceId = "72258d81-b535-4c3c-b286-d780ef90a847";
        const session = await client.authenticateDevice(deviceId, true, "mycustomusername");

        console.info("Successfully authenticated:", session);

        const account = await client.getAccount(session);
        const user = account.user;

        console.info("account:", account);

        await client.sessionLogout(session);

        /*
        var newUsername = "NotTheImp0ster";
        var newDisplayName = "Innocent Dave";
        var newLocation = "Edinburgh";

        var apiUpdateAccountRequest = new client.ApiUpdateAccountRequest();
        apiUpdateAccountRequest.newUsername = newUsername;
        apiUpdateAccountRequest.newDisplayName = newDisplayName;
        apiUpdateAccountRequest.newLocation = newLocation;

        await client.updateAccount(session, apiUpdateAccountRequest);

        var authToken = session.token;
        var refreshToken = session.refresh_token;

        session = session.restore(authToken, refreshToken);
        */
    } catch (e) {
        console.info(e);
    }
})();
