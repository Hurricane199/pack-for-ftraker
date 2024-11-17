package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	distance := distance(action)
	return distance / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	// ваш код здесь
	switch {
	case trainingType == "Бег":
		distance := distance(action)                               // вызовите здесь необходимую функцию
		speed := meanSpeed(action, duration)                       // вызовите здесь необходимую функцию
		calories := RunningSpentCalories(action, weight, duration) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Ходьба":
		distance := distance(action)                                       // вызовите здесь необходимую функцию
		speed := meanSpeed(action, duration)                               // вызовите здесь необходимую функцию
		calories := WalkingSpentCalories(action, duration, weight, height) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	case trainingType == "Плавание":
		distance := distance(action)                                               // вызовите здесь необходимую функцию
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)                // вызовите здесь необходимую функцию
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight) // вызовите здесь необходимую функцию
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	// ваш код здесь
	kA := float64(runningCaloriesMeanSpeedMultiplier) // сокращаю название константы [18.0]
	kB := runningCaloriesMeanSpeedShift               // сокращаю название константы [1.79]

	speed := meanSpeed(action, duration) // скорость в км/ч

	// ((18 * СредняяСкоростьВКм/ч * 1.79) * ВесСпортсменаВКг / mInKm  * ВремяТренировкиВЧасах * minInH)

	calories := ((kA * speed * kB) * weight / mInKm * duration * minInH)

	return calories
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	// ваш код здесь
	kA := walkingCaloriesWeightMultiplier // сокращаю название константы [0.035]
	kB := walkingSpeedHeightMultiplier    // сокращаю название константы  [0.029]

	speed := meanSpeed(action, duration) * kmhInMsec // скорость в м/с
	height /= 100                                    // перевожу рост из см в м

	// ((0.035 * ВесСпортсменаВКг + (СредняяСкоростьВМетрахВСекунду**2 / РостВМетрах) * 0.029 * ВесСпортсменаВКг) * ВремяТренировкиВЧасах * minInH)

	calories := ((kA*weight + (math.Pow(speed, 2)/height)*kB*weight) * duration * minInH)

	return calories
}

// Константы для расчета калорий, расходуемых при плавании.
const (
	swimmingCaloriesMeanSpeedShift   = 1.1 // среднее количество сжигаемых колорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier = 2   // множитель веса при плавании.
)

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	// ваш код здесь
	kA := swimmingCaloriesMeanSpeedShift            // сокращаю название константы [1.1]
	kB := float64(swimmingCaloriesWeightMultiplier) // сокращаю название константы [2.0]

	speed := swimmingMeanSpeed(lengthPool, countPool, duration) // скорость при плавании в км/ч

	// (СредняяСкоростьВКм/ч + 1.1) * 2 * ВесСпортсменаВКг * ВремяТренировкиВЧасах

	calories := (speed + kA) * kB * weight * duration
	return calories
}
