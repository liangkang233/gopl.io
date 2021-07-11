// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
/* Celsius 摄氏		Fahrenheit 华氏		Kelvin 开尔文		Feet 英尺		Meter 米
换算规则:
	华氏度 = 32＋摄氏度×1.8
	摄氏度 =（华氏度-32）÷1.8
	摄氏度 = K氏温度- 273.15
	英尺 = 米 / 3.2808 */

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Feet float64
type Meter float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	canshu                = 3.2808
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string       { return fmt.Sprintf("%gf", f) }

// 练习2.1 添加开尔文的转换规则
// 练习2.2 添加英尺与米的转换规则

//!-
