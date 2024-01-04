import { BrowserRouter, Routes, Route } from "react-router-dom"
import React, { Suspense } from "react"

const LandingPage = React.lazy(() => import("./pages/LandingPage"))


function App() {

  return(

    <BrowserRouter basename="/">

      <Suspense fallback={<div> Loading ..... </div>}>

        <Routes>
          <Route path="/" Component={LandingPage} />
        </Routes>

      </Suspense>

    </BrowserRouter>

  )
}

export default App
