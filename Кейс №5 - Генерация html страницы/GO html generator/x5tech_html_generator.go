package main

import (
	"fmt"
	"os"
	"time"
)

const htmlTemplate = `<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>X5 Tech | Наименование организации</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&amp;family=Roboto:wght@400;500;700&amp;family=Montserrat:wght@400;500;700&amp;family=Poppins:wght@400;500;600&amp;family=Space+Grotesk:wght@400;500;600&amp;display=swap">
    <style>
        @import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&amp;family=Roboto:wght@400;500;700&amp;family=Montserrat:wght@400;500;700&amp;family=Poppins:wght@400;500;600&amp;family=Space+Grotesk:wght@400;500;600&amp;display=swap');
        
        :root {
            --x5-red: #E30613;
        }
        
        body {
            font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
        }
        
        .section-header {
            font-family: 'Inter', system-ui, sans-serif;
            font-weight: 700;
            letter-spacing: -0.025em;
        }

        .x5-card {
            transition: transform 0.2s cubic-bezier(0.4, 0, 0.2, 1), box-shadow 0.2s;
        }
        
        .x5-card:hover {
            transform: translateY(-4px);
            box-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1);
        }

        .tech-badge {
            background: linear-gradient(135deg, #0F172A 0%%, #1E293B 100%%);
            color: white;
        }

        .font-sample {
            padding: 1rem;
            border-radius: 0.75rem;
            background: #F8FAFC;
            border: 1px solid #E2E8F0;
            margin-bottom: 1rem;
            transition: all 0.2s;
        }

        .font-label {
            font-size: 0.75rem;
            font-weight: 600;
            color: #64748B;
            margin-bottom: 0.5rem;
            text-transform: uppercase;
            letter-spacing: 0.05em;
        }

        .official-note {
            background: #FEF2F2;
            border-left: 4px solid #E30613;
        }
    </style>
</head>
<body class="bg-slate-50 text-slate-900">
    <!-- Header / Navbar -->
    <nav class="bg-white border-b border-slate-200 sticky top-0 z-50">
        <div class="max-w-screen-xl mx-auto px-6 py-4 flex items-center justify-between">
            <div class="flex items-center gap-x-3">
                <div class="w-10 h-10 bg-[#E30613] rounded-xl flex items-center justify-center">
                    <span class="text-white font-bold text-2xl tracking-tighter">X5</span>
                </div>
                <div>
                    <div class="font-bold text-2xl tracking-tighter">X5 Tech</div>
                </div>
            </div>
            
            <div class="hidden md:flex items-center gap-x-8 text-sm">
                <a href="#about" class="hover:text-[#E30613] transition-colors">О компании</a>
                <a href="#facts" class="hover:text-[#E30613] transition-colors">Факты и цифры</a>
                <a href="#products" class="hover:text-[#E30613] transition-colors">Продукты и решения</a>
                <a href="#fonts" class="hover:text-[#E30613] transition-colors font-medium">Варианты шрифтов</a>
            </div>
            
            <div class="flex items-center gap-x-3">
                <div class="px-4 py-1.5 bg-slate-100 rounded-full text-xs font-medium flex items-center gap-x-2">
                    <div class="w-2 h-2 bg-emerald-500 rounded-full animate-pulse"></div>
                    <span>Аккредитованная IT-компания</span>
                </div>
            </div>
        </div>
    </nav>

    <!-- Hero / Title -->
    <div class="max-w-screen-xl mx-auto px-6 pt-12 pb-8">
        <div class="max-w-3xl">

            
            <h1 class="text-6xl md:text-7xl font-bold tracking-tighter leading-none section-header">
                X5 Tech
            </h1>
            <p class="mt-4 text-2xl text-slate-600 font-medium">
                Общество с ограниченной ответственностью «ИТ ИКС 5 Технологии»
            </p>
            <p class="mt-3 text-xl text-slate-500 max-w-md">
                Основной цифровой партнёр торговых сетей и бизнесов X5
            </p>
            
            <div class="flex flex-wrap gap-3 mt-8">
                <a href="https://x5.tech" target="_blank" 
                   class="inline-flex items-center justify-center px-8 py-3.5 bg-[#E30613] hover:bg-red-700 transition-all text-white font-semibold rounded-2xl text-sm shadow-lg shadow-red-500/30">
                    Официальный сайт →
                </a>
                <a href="#about" 
                   class="inline-flex items-center justify-center px-8 py-3.5 border border-slate-300 hover:bg-white transition-all font-medium rounded-2xl text-sm">
                    Подробнее о компании
                </a>
            </div>
        </div>
    </div>

    <!-- About Section -->
    <div id="about" class="max-w-screen-xl mx-auto px-6 py-12 border-t border-slate-200">
        <div class="grid md:grid-cols-12 gap-x-12">
            <div class="md:col-span-5">
                <div class="sticky top-24">
                    <div class="uppercase tracking-[2px] text-xs font-semibold text-[#E30613] mb-3">КЕЙС-ЗАДАЧА № 5 • ГЕНЕРИРУЕМАЯ СТРАНИЦА</div>
                    <h2 class="text-4xl font-bold tracking-tighter">О компании</h2>
                    <p class="mt-4 text-slate-600">X5 Tech — IT-компания и основной цифровой партнёр X5. Мы разрабатываем решения, которые помогают миллионам покупателей и сотням тысяч сотрудников группы работать с максимальным технологическим комфортом.</p>
                    
                    <div class="mt-8 flex items-center gap-x-4 text-sm">
                        <div>
                            <div class="font-mono text-xs text-slate-500">ИНН</div>
                            <div class="font-semibold">1615014289</div>
                        </div>
                        <div class="h-3 w-px bg-slate-200"></div>
                        <div>
                            <div class="font-mono text-xs text-slate-500">ОКВЭД</div>
                            <div class="font-semibold">62.01</div>
                        </div>
                    </div>
                </div>
            </div>
            
            <div class="md:col-span-7 mt-10 md:mt-0">
                <div class="prose prose-slate max-w-none text-[15px]">
                    <p><strong>Организационно-правовая форма:</strong> Общество с ограниченной ответственностью (ООО).</p>
                    
                    <p class="mt-4"><strong>Основные виды деятельности:</strong></p>
                    <ul class="list-disc pl-5 space-y-1.5 text-slate-600">
                        <li>Разработка компьютерного программного обеспечения и информационных систем</li>
                        <li>Полный цикл разработки цифровых продуктов: от анализа требований и архитектуры до внедрения и поддержки</li>
                        <li>Разработка и сопровождение мобильных приложений, веб-сервисов, платформ больших данных и облачных решений</li>
                        <li>Внедрение технологий искусственного интеллекта, машинного обучения, видеоаналитики и автоматизации процессов ритейла</li>
                        <li>Обеспечение работы более 25 000 магазинов и 52 распределительных центров X5</li>
                    </ul>
                    
                    <p class="mt-6"><strong>Территориальное размещение:</strong></p>
                    <p class="text-slate-600">Юридический адрес: 420500, Республика Татарстан, г. Иннополис, ул. Университетская, д. 7, офис 403 (резидент ОЭЗ «Иннополис»).</p>
                    <p class="text-slate-600 mt-1">Операционные офисы и команды: Москва (основные офисы «Калитники» и «Оазис»), Санкт-Петербург, Ижевск, Нижний Новгород. Возможна полностью удалённая работа.</p>
                </div>
            </div>
        </div>
    </div>

    <!-- Key Facts -->
    <div id="facts" class="bg-white border-y border-slate-200 py-12">
        <div class="max-w-screen-xl mx-auto px-6">
            <h3 class="font-semibold tracking-tight text-2xl mb-8">Ключевые факты и цифры (2024–2025)</h3>
            
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                <div class="x5-card bg-slate-900 text-white p-6 rounded-3xl">
                    <div class="text-5xl font-bold tracking-tighter">6000+</div>
                    <div class="mt-2 text-sm opacity-75">специалистов в команде</div>
                    <div class="text-[10px] mt-4 opacity-60">80%% — мидлы и сеньоры</div>
                </div>
                <div class="x5-card bg-white border border-slate-200 p-6 rounded-3xl">
                    <div class="text-5xl font-bold tracking-tighter text-slate-900">12+ ПБ</div>
                    <div class="mt-2 text-sm text-slate-600">объём кластера больших данных</div>
                    <div class="text-[10px] mt-4 text-slate-500">&gt;100 млрд событий в день</div>
                </div>
                <div class="x5-card bg-white border border-slate-200 p-6 rounded-3xl">
                    <div class="text-5xl font-bold tracking-tighter text-slate-900">600+</div>
                    <div class="mt-2 text-sm text-slate-600">информационных систем в эксплуатации</div>
                    <div class="text-[10px] mt-4 text-slate-500">2600+ физических серверов</div>
                </div>
                <div class="x5-card bg-white border border-slate-200 p-6 rounded-3xl">
                    <div class="text-5xl font-bold tracking-tighter text-slate-900">70+</div>
                    <div class="mt-2 text-sm text-slate-600">цифровых продуктов</div>
                    <div class="text-[10px] mt-4 text-slate-500">400+ проектов в работе</div>
                </div>
            </div>
        </div>
    </div>

    <!-- Products & Tech -->
    <div id="products" class="max-w-screen-xl mx-auto px-6 py-12">
        <h3 class="font-semibold tracking-tight text-2xl mb-6">Основные продукты и технологические решения</h3>
        
        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div class="x5-card p-5 border border-slate-200 rounded-2xl bg-white">
                <div class="font-semibold">Мобильные приложения X5</div>
                <div class="text-sm text-slate-600 mt-1.5">Приложения для «Пятёрочки», «Перекрёстка», «Чижика», «Много лосося» и 5Post. Push-уведомления, персонализация, геймификация.</div>
            </div>
            <div class="x5-card p-5 border border-slate-200 rounded-2xl bg-white">
                <div class="font-semibold">X5 Salt — собственное облако</div>
                <div class="text-sm text-slate-600 mt-1.5">Open-source облачная платформа на базе OpenStack / Kubernetes. Высокий уровень безопасности и независимости.</div>
            </div>
            <div class="x5-card p-5 border border-slate-200 rounded-2xl bg-white">
                <div class="font-semibold">Dialog X5</div>
                <div class="text-sm text-slate-600 mt-1.5">Платформа для взаимодействия с поставщиками и партнёрами (8 цифровых продуктов).</div>
            </div>
            <div class="x5-card p-5 border border-slate-200 rounded-2xl bg-white">
                <div class="font-semibold">Nexus WMS + TMS</div>
                <div class="text-sm text-slate-600 mt-1.5">Системы управления складами и транспортом (автоматизация 52 РЦ и собственного автопарка &gt;4000 грузовиков).</div>
            </div>
            <div class="x5-card p-5 border border-slate-200 rounded-2xl bg-white">
                <div class="font-semibold">AI / ML решения</div>
                <div class="text-sm text-slate-600 mt-1.5">Прогнозирование спроса, динамическое ценообразование, видеоаналитика, A/B-тестирование, персонализация ассортимента.</div>
            </div>
            <div class="x5-card p-5 border border-slate-200 rounded-2xl bg-white">
                <div class="font-semibold">Кассы самообслуживания (КСО)</div>
                <div class="text-sm text-slate-600 mt-1.5">Собственная разработка, машинное зрение и интеграция с экосистемой.</div>
            </div>
        </div>
        
        <div class="mt-8 text-xs text-slate-500 flex items-center gap-x-2">
            <span>Технологический стек:</span> 
            <span class="font-mono text-[12px] bg-slate-100 px-2 py-0.5 rounded">Java • Spring Boot • Kotlin • Python • Go • React • Kubernetes • Kafka • Greenplum • Airflow • CatBoost • Prometheus • Grafana и др.</span>
        </div>
    </div>

    <!-- Fonts Section -->
    <div id="fonts" class="bg-white border-t border-slate-200 py-12">
        <div class="max-w-screen-xl mx-auto px-6">
            <div class="max-w-2xl">
                <div class="uppercase text-xs tracking-[1.5px] font-semibold text-[#E30613]">ДИЗАЙН • ТИПОГРАФИКА</div>
                <h3 class="text-3xl font-bold tracking-tighter mt-2">Официальный шрифт и 5 вариантов для стилизации</h3>
                <p class="mt-3 text-slate-600">Официальный фирменный шрифт X5 — <strong>X5 Sans</strong> (и X5 Sans UI для интерфейсов). Это современный лаконичный геометрический гротеск, разработанный специально для X5 совместно с ParaType. В веб-версии мы приближаем его с помощью Inter и аналогичных шрифтов.</p>
            </div>
            
            <div class="mt-10">
                <div class="font-sample">
                    <div class="font-label flex items-center justify-between">
                        <span>Вариант 1 — Inter</span>
                        <span class="font-mono text-emerald-600 text-xs">font-family: 'Inter', system-ui, sans-serif;</span>
                    </div>
                    <p class="text-xl font-medium" style="font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;">
                        X5 Tech — IT-компания и основной цифровой партнёр торговых сетей и бизнесов X5. Мы помогаем миллионам покупателей покупать свежие продукты быстро и удобно.
                    </p>
                </div>
                
                <div class="font-sample">
                    <div class="font-label flex items-center justify-between">
                        <span>Вариант 2 — Roboto</span>
                        <span class="font-mono text-emerald-600 text-xs">font-family: 'Roboto', system-ui, sans-serif;</span>
                    </div>
                    <p class="text-xl font-medium" style="font-family: 'Roboto', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;">
                        X5 Tech — IT-компания и основной цифровой партнёр торговых сетей и бизнесов X5. Мы помогаем миллионам покупателей покупать свежие продукты быстро и удобно.
                    </p>
                </div>
                
                <div class="font-sample">
                    <div class="font-label flex items-center justify-between">
                        <span>Вариант 3 — Montserrat</span>
                        <span class="font-mono text-emerald-600 text-xs">font-family: 'Montserrat', system-ui, sans-serif;</span>
                    </div>
                    <p class="text-xl font-medium" style="font-family: 'Montserrat', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;">
                        X5 Tech — IT-компания и основной цифровой партнёр торговых сетей и бизнесов X5. Мы помогаем миллионам покупателей покупать свежие продукты быстро и удобно.
                    </p>
                </div>
                
                <div class="font-sample">
                    <div class="font-label flex items-center justify-between">
                        <span>Вариант 4 — Poppins</span>
                        <span class="font-mono text-emerald-600 text-xs">font-family: 'Poppins', system-ui, sans-serif;</span>
                    </div>
                    <p class="text-xl font-medium" style="font-family: 'Poppins', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;">
                        X5 Tech — IT-компания и основной цифровой партнёр торговых сетей и бизнесов X5. Мы помогаем миллионам покупателей покупать свежие продукты быстро и удобно.
                    </p>
                </div>
                
                <div class="font-sample">
                    <div class="font-label flex items-center justify-between">
                        <span>Вариант 5 — Space Grotesk</span>
                        <span class="font-mono text-emerald-600 text-xs">font-family: 'Space Grotesk', system-ui, sans-serif;</span>
                    </div>
                    <p class="text-xl font-medium" style="font-family: 'Space Grotesk', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;">
                        X5 Tech — IT-компания и основной цифровой партнёр торговых сетей и бизнесов X5. Мы помогаем миллионам покупателей покупать свежие продукты быстро и удобно.
                    </p>
                </div>
            </div>
        </div>
    </div>

        <div class="max-w-screen-xl mx-auto px-6 text-center text-xs text-slate-500">
            <p>Страница сгенерирована в рамках индивидуального задания на производственную (технологическую) практику</p>
            <p class="mt-1">Данные основаны на общедоступной информации с официального сайта <a href="https://x5.tech" class="underline hover:text-slate-700">x5.tech</a> и открытых источников. Актуально на %s</p>
        </div>
    <script>
        function initializeTailwind() {
            console.log('%%c[X5 Tech Practice Page] Tailwind + custom fonts initialized', 'color:#64748b');
        }
        window.onload = initializeTailwind;
    </script>
</body>
</html>`

func main() {
	currentDate := time.Now().Format("02.01.2006")
	output := fmt.Sprintf(htmlTemplate, currentDate)

	filename := "x5_tech_generated.html"
	err := os.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return
	}

	fmt.Printf("HTML-страница успешно сгенерирована: %s\n", filename)
	fmt.Println("Открыть файл в браузере и проверить результат.")
}
