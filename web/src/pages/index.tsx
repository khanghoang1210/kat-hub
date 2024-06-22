import Image from "next/image";
import { Inter } from "next/font/google";
import Topbar from "@/components/Topbar/Topbar";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    <div className="bg-blue-900">
      <Topbar/>
    </div>
  );
}
