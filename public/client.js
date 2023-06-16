class ShareMe {
    static VERSION = "0.4.2";
    _server = "";
    constructor(server) {
        this._server = server;
    }
    /** helper method for to build url with namespace  */
    apiUrl(namespace) {
        const url = new URL(this._server);
        url.pathname = `/${namespace}`;
        return url;
    }

    // those are refs to function that can cancel previous request
    cancelGetFunc = () => {};
    cancelSetFunc = () => {};

    /**
     * @returns {Promise<string|null>}
     */
    async get(namespace) {
        if (!isNamespaceValid(namespace)) {
            console.error("Invalid namespace", namespace);
            return null; // fail to get as namespace is invalid
        }
        try {
            this.cancelGetFunc(); // cancel previous get request
            const { res, cancel } = cancelablePost(this.apiUrl(namespace));
            this.cancelGetFunc = cancel; // replace with new cancel function

            const resp = await res;
            if (!resp.ok) return null; // return null means error

            return await resp.text(); // return text
        } catch {
            return null;
        }
    }

    /**
     * @returns {Promise<boolean>}
     */
    async set(namespace, t) {
        if (!isNamespaceValid(namespace)) {
            console.error("Invalid namespace", namespace);
            return false; // fail to set as namespace is invalid
        }

        try {
            this.cancelSetFunc(); // cancel previous set request
            const { res, cancel } = cancelablePost(this.apiUrl(namespace), JSON.stringify({ t }));
            this.cancelSetFunc = cancel; // replace with new cancel function

            const resp = await res;
            if (!resp.ok) return false; // fail to set

            return true; // success
        } catch {
            return false;
        }
    }
}

function cancelablePost(url, body = undefined) {
    const ab = new AbortController();
    const res = fetch(url, {
        headers: body === undefined ? undefined : { "content-type": "application/json" },
        method: "POST",
        body,
        signal: ab.signal,
    });
    return { cancel: () => ab.abort(), res };
}

function isNamespaceValid(namespace) {
    return /^[a-zA-Z0-9]{1,16}$/.test(namespace);
}

export { ShareMe, isNamespaceValid };
