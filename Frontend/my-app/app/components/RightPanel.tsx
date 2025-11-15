import React from "react";
import { PanelContext } from "../context/PanelContext";

function Wydatki() {
  return (
    <div className="p-6 h-full">
      <h3>Wydatki</h3>
    </div>
  );
}

function Przychody() {
  return (
    <div className="p-6 h-full">
      <h3>Przychody</h3>
    </div>
  );
}

function Oszczednosci() {
  return (
    <div className="p-6 h-full">
      <h3>Oszczędności</h3>
    </div>
  );
}

function Obrazy() {
  return (
    <div className="p-6 h-full">
      <h3>Obrazy</h3>
    </div>
  );
}

function Profil() {
  return (
    <div className="p-6 h-full">
    <h3>Profil</h3>
    </div>
  );
}
export function RightPanel(): React.ReactElement {
  const { selected } = React.useContext(PanelContext);

  return (
    <div className="h-[583px] w-[986px] bg-[#C0C3C2] rounded-2xl overflow-hidden mt-5">
    
      <div className="w-full h-[97px] bg-[#1A6558]  text-white px-4 py-2 border-b border-white/10">
        <h2 className="text-[40px] font-semibold capitalize">
          {selected}
        </h2>
      </div>


      
      {selected === "wydatki" && <Wydatki />}
      {selected === "przychody" && <Przychody />}
      {selected === "oszczednosci" && <Oszczednosci />}
      {selected === "obrazy" && <Obrazy />}
      {selected === "profil" && <Profil />}
    </div>
  );
}



export default RightPanel;
