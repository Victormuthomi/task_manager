import requests
from django.http import JsonResponse
from django.views import View
from geopy.distance import geodesic

# Constants
MAX_RANGE = 500  # Max range in miles per full tank
MPG = 10  # Miles per gallon

# Load fuel stations (Assume this is coming from a database or static JSON file)
FUEL_STATIONS = [
    {"name": "Station A", "lat": 40.7128, "lon": -74.0060, "price": 3.50},
    {"name": "Station B", "lat": 39.9526, "lon": -75.1652, "price": 3.45},
    {"name": "Station C", "lat": 38.8977, "lon": -77.0365, "price": 3.40},
]

# Free Routing API (Replace with your API Key)
ROUTING_API_URL = "https://router.project-osrm.org/route/v1/driving/{},{};{},{}?overview=full"

class RouteView(View):
    def get(self, request):
        start = request.GET.get("start")  # e.g., "New York, NY"
        finish = request.GET.get("finish")  # e.g., "Los Angeles, CA"
        
        if not start or not finish:
            return JsonResponse({"error": "Start and Finish locations are required."}, status=400)
        
        # Convert locations to coordinates
        start_coords = self.get_coordinates(start)
        finish_coords = self.get_coordinates(finish)
        
        if not start_coords or not finish_coords:
            return JsonResponse({"error": "Invalid start or finish location."}, status=400)
        
        # Get route distance
        distance, route_geometry = self.get_route(start_coords, finish_coords)
        if distance is None:
            return JsonResponse({"error": "Could not calculate route."}, status=500)
        
        # Get fuel stops
        fuel_stops, total_fuel_cost = self.get_fuel_stops(start_coords, finish_coords, distance)
        
        return JsonResponse({
            "start": start,
            "finish": finish,
            "distance_miles": round(distance, 2),
            "fuel_stops": fuel_stops,
            "total_fuel_cost": round(total_fuel_cost, 2),
            "route_geometry": route_geometry
        })
    
    def get_coordinates(self, location):
        """ Convert location name to coordinates using OpenStreetMap """
        response = requests.get(f"https://nominatim.openstreetmap.org/search?q={location}&format=json")
        data = response.json()
        if data:
            return float(data[0]["lat"]), float(data[0]["lon"])
        return None
    
    def get_route(self, start_coords, finish_coords):
        """ Fetches route and distance using OSRM API """
        url = ROUTING_API_URL.format(start_coords[1], start_coords[0], finish_coords[1], finish_coords[0])
        response = requests.get(url)
        data = response.json()
        if "routes" in data and data["routes"]:
            return data["routes"][0]["distance"] * 0.000621371, data["routes"][0]["geometry"]  # Convert meters to miles
        return None, None
    
    def get_fuel_stops(self, start_coords, finish_coords, distance):
        """ Calculates fuel stops and cost """
        stops_needed = max(0, int(distance // MAX_RANGE))
        fuel_stops = []
        total_fuel_cost = 0
        
        for i in range(stops_needed):
            # Estimate stop location (every 500 miles along route)
            stop_distance = (i + 1) * MAX_RANGE
            stop_coords = self.estimate_stop_location(start_coords, finish_coords, stop_distance, distance)
            
            # Find the cheapest station nearby
            cheapest_station = min(
                FUEL_STATIONS,
                key=lambda s: geodesic((s["lat"], s["lon"]), stop_coords).miles if geodesic((s["lat"], s["lon"]), stop_coords).miles < 50 else float("inf"),
            )
            
            if cheapest_station["price"] != float("inf"):
                fuel_stops.append(cheapest_station)
                total_fuel_cost += (MAX_RANGE / MPG) * cheapest_station["price"]
        
        return fuel_stops, total_fuel_cost
    
    def estimate_stop_location(self, start_coords, finish_coords, stop_distance, total_distance):
        """ Approximates stop coordinates based on distance along route """
        lat1, lon1 = start_coords
        lat2, lon2 = finish_coords
        
        fraction = stop_distance / total_distance
        stop_lat = lat1 + (lat2 - lat1) * fraction
        stop_lon = lon1 + (lon2 - lon1) * fraction
        
        return stop_lat, stop_lon

