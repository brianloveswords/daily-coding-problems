import scala.util.Try
import scala.util.Success
import scala.util.Failure
import scala.util.control._

class NumberOutOfBounds(val n: Int) extends Exception {
  override def getMessage(): String =
    s"Number must be between 0 and 59271, got ${n}"
}
class MissingCLIArgument extends Exception {
  override def getMessage(): String = s"Expected at least 1 argument"
}

object Main extends App {
  val MAX_SEQUENCE_VALUE = 59271
  def eprintln(x: Any) = Console.err.println(x)

  tryIndex(args, 0)
    .flatMap(tryStrToInt)
    .flatMap(validateRange)
    .map(getPerfectNumber) match {
    case Failure(exception) =>
      eprintln(
        s"[${Console.RED}ERROR${Console.RESET}]: ${exception.getMessage()}"
      )
    case Success(value) => println(value)
  }

  //---

  def tryIndex[T](itr: Array[T], idx: Int): Try[T] =
    try { Success(itr(idx)) }
    catch { case e: Exception => Failure(new MissingCLIArgument) }

  def tryStrToInt(str: String): Try[Int] = Try { str.toInt }

  def validateRange(n: Int): Try[Int] =
    n match {
      case n if n <= 0 || n > MAX_SEQUENCE_VALUE =>
        Failure(new NumberOutOfBounds(n))
      case _ => Success(n)
    }

  def getPerfectNumber(sequenceId: Int): Int = {
    var i = 0;
    var current = 19;
    var nextPower = 100;

    while (true) {
      if (isPerfect(current)) {
        i += 1
      }
      if (current == nextPower) {
        eprintln(
          s"[${Console.MAGENTA}DEBUG${Console.RESET}]: passed ${nextPower}"
        );
        nextPower *= 10;
      }
      if (i == sequenceId) {
        return current
      }
      current += 1
    }
    current
  }

  def isPerfect(n: Int): Boolean = {
    var current = n;
    var sum = 0;
    while (true) {
      var remaining_places = current / 10
      if (remaining_places == 0) {
        sum += current
        return sum == 10
      }
      sum += current - remaining_places * 10
      current = remaining_places
    }
    return false
  }

}
