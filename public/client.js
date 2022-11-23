class ShareMe {
    static VERSION = "0.4.1";
    constructor(server) {
        this.server = server;
    }
    async get(namespace) {
        // => string|null
        if (!isNamespaceValid(namespace)) throw new Error("Invalid namespace");
        try {
            const url = new URL(this.server);
            url.pathname = `/${namespace}`;
            const resp = await fetch(url.toString(), {
                method: "POST",
                // use POST with no body to get data
            });
            if (!resp.ok) return null;
            return await resp.text();
        } catch (_e) {
            return null;
        }
    }

    async set(namespace, t) {
        // => boolean
        if (!isNamespaceValid(namespace)) throw new Error("Invalid namespace");
        try {
            const url = new URL(this.server);
            url.pathname = `/${namespace}`;
            const resp = await fetch(url.toString(), {
                method: "POST",
                body: JSON.stringify({ t }),
                headers: {
                    "content-type": "application/json",
                },
            });
            if (!resp.ok) return false;
            return true;
        } catch (_e) {
            return false;
        }
    }
}

function isNamespaceValid(namespace) {
    return /^[a-zA-Z0-9]{1,16}$/.test(namespace);
}

export { ShareMe, isNamespaceValid };
