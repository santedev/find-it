const user = {
  name: "",
  id: "",
  email: "",
  picture: "",
};

const OAuthGoogleData = {
  code: "",
  state: "",
  status: false
};

export function getSession() {
  return { ...user, ...OAuthGoogleData };
}

export function getOAuthGoogleData() {
  return OAuthGoogleData;
}

export function getUser() {
  return user;
}

export function setUserValue(key, value) {
  try {
    if (user[key] == null) {
      throw new Error(
        "key in setUserValue doesnt exist or value is null or undefined"
      );
    }
    user[key] = value;
  } catch (error) {
    console.error(error);
  }
}

export function setOAuthGoogleValue(key, value) {
  try {
    if (OAuthGoogleData[key] == null) {
      throw new Error(
        "key in setOAuthGoogleValue doesnt exist or value is null or undefined"
      );
    }
    OAuthGoogleData[key] = value;
  } catch (error) {
    console.error(error);
  }
}
