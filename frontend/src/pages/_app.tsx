import type { AppProps } from "next/app";
import '@/style/globals.css'
import NavBar from "@/components/NavBar";
import Footer from "@/components/Footer";

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <NavBar />
      <Component {...pageProps} />
      <Footer />
    </>
  );
}
