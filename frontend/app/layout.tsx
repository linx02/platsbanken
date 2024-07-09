import type { Metadata } from "next";
import { Noto_Sans } from "next/font/google";
import "./globals.css";
import Nav from "./components/Nav";

const notoSans = Noto_Sans({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Platsbanken - fast bättre",
  description: "Sök jobb i platsbanken",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={notoSans.className}>
        <Nav />
        {children}</body>
    </html>
  );
}
