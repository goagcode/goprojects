'use strict'

const express = require('express')
const app = express()

const path = require('path')
const PROTO_PATH = path.join('pb', 'messages.proto')
const SERVER_ADDR = 'localhost:3030'
const grpc = require('grpc')

function helloHandler(name, callback) {
  const HelloService = grpc.load(PROTO_PATH).HelloService
  const client = new HelloService(SERVER_ADDR, grpc.credentials.createInsecure())

  client.sayHello({name: name}, (error, res) => {
    if (error) {
      console.log(error)
      return
    }
    console.log('I got a message => ', res.message)
    return callback(null, res)
  })
}

app.get('/say/:name', (req, res) => {
  helloHandler(req.params.name, (err, message) => {
    res.status(200).json(message)
  })
})

app.listen(8080, () => {
  console.log('server running at localhost:8080')
})
