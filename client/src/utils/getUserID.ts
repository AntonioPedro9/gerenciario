import Cookies from "js-cookie";
import jwt_decode from "jwt-decode";

type DecodedToken = {
  sub: string;
};

export default function getUserID() {
  const token = Cookies.get("Authorization");

  if (token !== undefined) {
    const decoded = jwt_decode(token) as DecodedToken;
    const userID = decoded.sub;

    return userID;
  }
}
