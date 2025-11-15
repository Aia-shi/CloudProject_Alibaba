import React from "react";
import { PanelContext } from "../context/PanelContext";

export default function Navbar(): React.ReactElement {
  const { selected, setSelected } = React.useContext(PanelContext);

  return (
    <div className="w-[238px] h-[398px] rounded-r-2xl  flex flex-col gap-3  justify-center items-start"
      style={{
        background: "#1A6458"
      }}>

      <button
        className={`transition-all duration-200 origin-left font-semibold rounded-r-2xl text-white  ${
          selected === "wydatki" 
            ? "w-[269px] h-[47px] opacity-100" 
            : "w-[212px] h-[47px] opacity-70"
        }`}
        style={{ backgroundColor: "#D9D9D9", color: "#000" }}
        onClick={() => setSelected("wydatki")}
      >
        Wydatki
      </button>

      <button
        className={`transition-all duration-200 font-semibold rounded-r-2xl text-white ${
          selected === "przychody" 
            ? "w-[269px] h-[47px] opacity-100" 
            : "w-[212px] h-[47px] opacity-70"
        }`}
        style={{ backgroundColor: "#D9D9D9", color: "#000" }}
        onClick={() => setSelected("przychody")}
      >
        Przychody
      </button>

      <button
        className={`transition-all duration-200 font-semibold rounded-r-2xl text-white ${
          selected === "oszczednosci" 
            ? "w-[269px] h-[47px] opacity-100" 
            : "w-[212px] h-[47px] opacity-70"
        }`}
        style={{ backgroundColor: "#D9D9D9", color: "#000" }}
        onClick={() => setSelected("oszczednosci")}
      >
        Oszczędności
      </button>
        
      <button
        className={`transition-all duration-200 font-semibold rounded-r-2xl text-white ${
          selected === "obrazy" 
            ? "w-[269px] h-[47px] opacity-100" 
            : "w-[212px] h-[47px] opacity-70"
        }`}
        style={{ backgroundColor: "#D9D9D9", color: "#000" }}
        onClick={() => setSelected("obrazy")}
      >
        Obrazy

      </button>
            <button
        className={`transition-all duration-200 font-semibold rounded-r-2xl text-white ${
          selected === "profil" 
            ? "w-[269px] h-[47px] opacity-100" 
            : "w-[212px] h-[47px] opacity-70"
        }`}
        style={{ backgroundColor: "#D9D9D9", color: "#000" }}
        onClick={() => setSelected("profil")}
      >
        Profil
      </button>
    </div>
  );
}

