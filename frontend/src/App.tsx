import React from 'react'
import  Sidebar  from "./Components/Sidebar"
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Bmain from './Booking/Bmain'
import Bsub from './Booking/Bsub'
import Confirm from './Booking/Confirm'
import List from './Booking/List'
import LogOut from './Page/LogOut'
import Student from './Page/Student'
import Report from './Page/Report'
import Form1 from './Page/Form1'
import Form2 from './Page/Form2'
import Form3 from './Page/Form3'
import Status from './Page/Status'

function App()  {
  return (
    <BrowserRouter>
      <Sidebar>
          <Routes>
            <Route path='/' element={<Bmain />} />
            <Route path='/Bmain' element={<Bmain/>} />
            <Route path='/Bsub' element={<Bsub/>} />
            <Route path='/Confirm' element={<Confirm/>} />
            <Route path='/List' element={<List/>} />
            <Route path='/LogOut' element={<LogOut/>} />
            <Route path='/Student' element={<Student/>} />
            <Route path='/Report' element={<Report/>} />
            <Route path='/Form1' element={<Form1/>} />
            <Route path='/Form2' element={<Form2/>} />
            <Route path='/Form3' element={<Form3/>} />
            <Route path='/Status' element={<Status/>} />
          </Routes>
      </Sidebar>
    </BrowserRouter>
  )
}

export default App