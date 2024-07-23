export function initSession() {
  if (import.meta.client) {
    const url = new URLSearchParams(window.location.search);
    const stateFromURL = url.get("state") ?? "";
    const codeFromURL = url.get("code") ?? "";

    setOAuthGoogleValue("state", stateFromURL);
    setOAuthGoogleValue("code", codeFromURL);

    const cleanUrl = window.location.origin + window.location.pathname;
    history.replaceState(null, "", cleanUrl);

    const sessionData = getSession();
    const ggCode = sessionData.code;
    const state = sessionStorage.getItem("state") || "stateFailed";
    const checkState = state === sessionData.state;
    const reqBody = {
      googleCode: ggCode,
    };
    if (ggCode.length > 0 && checkState) {
      (async function getSessionGoogle() {
        try {
          const response = await fetch("http://localhost:8000/auth/google", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify(reqBody),
          });
          const data = await response.text();
          if (response.ok) {
            const user = JSON.parse(data);
            if (!user) {
              throw new Error("fetched data is undefined or null");
            }
            setUserValue("name", user.name);
            setUserValue("email", user.email);
            setUserValue("id", user.user_id);
            setUserValue("picture", user.picture);
            setOAuthGoogleValue("status", true);
            console.log(getSession());
          } else {
            console.error("Error:", data);
          }
        } catch (error) {
          console.log(error);
        }
      })();
    } else {
      (async function getSessionByCookies() {
        try {
          const response = await fetch("http://localhost:8000/session/init", {
            method: "GET",
            credentials: "include",
          });
          const data = await response.text();
          if (response.ok) {
            const user = JSON.parse(data);
            if (!user) {
              throw new Error("fetched data is undefined or null");
            }
            setUserValue("name", user.name);
            setUserValue("email", user.email);
            setUserValue("id", user.user_id);
            setUserValue("picture", user.picture);
            setOAuthGoogleValue("status", true);
            console.log(getSession());
          } else {
            console.log(data);
          }
        } catch (error) {
          console.error(error);
        }
      })();
    }
  }
}

export function oauthSignIn() {
  if (import.meta.client) {
    try {
      async function init() {
        try {
          const response = await fetch(
            "http://localhost:8000/session/getClient",
            {
              method: "GET",
              credentials: "include",
            }
          );
          const data = await response.text();
          if (response.ok) {
            if (response.status === 200) return data;
          } else {
            console.log(data);
          }
        } catch (error) {
          console.error(error);
        }
        return;
      }
      (async () => {
        let client_id = await init();
        if (typeof client_id !== "string" || client_id === "") {
          throw new Error("Invalid or empty client ID received");
        }
        // Google's OAuth 2.0 endpoint for requesting an access token
        const oauth2Endpoint = "https://accounts.google.com/o/oauth2/v2/auth";

        // Create <form> element to submit parameters to OAuth 2.0 endpoint.
        const form = document.createElement("form");
        form.setAttribute("method", "GET"); // Send as a GET request.
        form.setAttribute("action", oauth2Endpoint);

        const state = Math.random().toString(32).substring(2);
        sessionStorage.setItem("state", state);
        // Parameters to pass to OAuth 2.0 endpoint.
        const params = {
          client_id: client_id || "",
          redirect_uri: "http://localhost:3000",
          access_type: "offline",
          response_type: "code",
          scope:
            "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email",
          include_granted_scopes: "true",
          state: state,
        };

        // Add form parameters as hidden input values.
        for (const p in params) {
          const input = document.createElement("input");
          input.setAttribute("type", "hidden");
          input.setAttribute("name", p);
          input.setAttribute("value", params[p]);
          form.appendChild(input);
        }
        document.body.appendChild(form);
        form.submit();
      })();
    } catch (error) {
      console.error(error);
    }
  }
}
