import "./globals.css";

import type { Metadata } from "next";
import { PlusJakartaSans } from "./fonts";
import { tw } from "@/libs/common";
import Providers from "./providers";

export const metadata: Metadata = {
  title: {
    default: "SpaceNotes",
    template: "%s | SpaceNotes",
  },
  description: "Powered by - Next.js",
};

export default function RootLayout(props: React.PropsWithChildren) {
  return (
    <html
      lang="en"
      className={tw("scroll-pt-20", PlusJakartaSans.variable)}
      suppressHydrationWarning
    >
      <body>
        <Providers>{props.children}</Providers>
      </body>
    </html>
  );
}
