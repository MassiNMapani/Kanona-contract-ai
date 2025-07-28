import "../styles/globals.css";
import { RoleProvider } from "../context/RoleContext";

export default function App({ Component, pageProps }) {
  return (
    <RoleProvider>
      <Component {...pageProps} />
    </RoleProvider>
  );
}
