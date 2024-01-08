import { BrowserRouter, Routes, Route } from "react-router-dom"
import React, { Suspense } from "react";

const LandingPage = React.lazy(() => import("./pages/LandingPage"))
const VaultPage = React.lazy(() => import("./pages/VaultPage"))

function App() {

  return (
    <BrowserRouter basename="/">
      <Suspense fallback={<div> Loading ..... </div>}>
        <Routes>

          <Route path="/" Component={LandingPage} />
          <Route path="/my-vault" Component={VaultPage}/>

        </Routes>
      </Suspense>
    </BrowserRouter>
  );
}


export default App
