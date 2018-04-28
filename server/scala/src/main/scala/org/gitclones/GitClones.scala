package org.gitclones

import slick.driver.H2Driver.api._
import scala.concurrent.ExecutionContext.Implicits.global
import org.gitclones.persistence

object GitClonesApp extends App {
  var db = Database.forConfig("h2mem1")
  try {
    println("test")
  } finally db.close
}
