package org.gitclones.persistence

//import scala.slick.driver.H2Driver.simple._
import slick.driver.H2Driver.api._
//import org.gitclones.persistence.GitClones._
import org.gitclones.model.GitClone

class GitClones(tag: Tag) extends Table[GitClone](tag, "GITCLONES") {
  def name = column[String]("NAME", O.Length(512))
  def url = column[String]("URL", O.PrimaryKey, O.Length(512))
  def * = (name, url) <> (GitClone.tupled, GitClone.unapply)
}
